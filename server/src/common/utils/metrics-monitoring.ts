'use strict';
const opentelemetry = require('@opentelemetry/sdk-node');
const { HttpInstrumentation } = require('@opentelemetry/instrumentation-http');
const {
  AwsInstrumentation,
} = require('@opentelemetry/instrumentation-aws-sdk');
const {
  MongooseInstrumentation,
} = require('@opentelemetry/instrumentation-mongoose');
const {
  OTLPTraceExporter,
} = require('@opentelemetry/exporter-trace-otlp-http');
const { Resource } = require('@opentelemetry/resources');
const {
  SemanticResourceAttributes,
} = require('@opentelemetry/semantic-conventions');
const {
  ExpressInstrumentation,
} = require('@opentelemetry/instrumentation-express');
const {
  KafkaJsInstrumentation,
} = require('opentelemetry-instrumentation-kafkajs');
import { Span } from '@opentelemetry/api';
import { getNodeAutoInstrumentations } from '@opentelemetry/auto-instrumentations-node';
import { v4 as uuidv4 } from 'uuid';

const exporterOptions = {
  url: process.env.SIGNOZ_HOST || 'http://172.17.0.1:4318/v1/traces',
};

const instrumentations = [
  getNodeAutoInstrumentations(),
  new ExpressInstrumentation({
    ignoreLayersType: ['middleware', 'request_handler'],
  }),
  new HttpInstrumentation(),
  new MongooseInstrumentation(),
  new AwsInstrumentation(),
  new KafkaJsInstrumentation({
    consumerHook: (span: Span, topic, message) => {
      span.setAttribute(
        'correlationId',
        message?.headers['x-correlation-id']?.toString() || uuidv4(),
      );
    },
  }),
];

const sdkOptions = {
  traceExporter: new OTLPTraceExporter(exporterOptions),
  instrumentations: instrumentations,
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]:
      process.env.NODE_ENV +
      '-' +
      (process.env.SERVICE_NAME || 'x2-core-algo-risk-analysis-backend'),
  }),
};

const sdk = new opentelemetry.NodeSDK(sdkOptions);

process.on('SIGTERM', () => {
  sdk
    .shutdown()
    .then(() => console.log('Tracing terminated'))
    .catch((error) => console.log('Error terminating tracing', error))
    .finally(() => process.exit(0));
});

export default sdk;
