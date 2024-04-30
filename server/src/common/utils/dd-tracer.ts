import tracer from 'dd-trace';

tracer.init({
  runtimeMetrics: true,
}); // initialized in a different file to avoid hoisting.

// Express server
tracer.use('express', {
  headers: ['x-correlation-id'],
  middleware: false,
});

// HTTP clients
tracer.use('http', {
  client: {
    headers: ['x-correlation-id'],
  },
});

// KafkaJS Nest Microservice
tracer.use('kafkajs');

export default tracer;
