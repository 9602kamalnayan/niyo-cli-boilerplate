import {
    LoggerService,
    ILoggingMetaDataProvider,
    LogLevel,
} from 'nc-node-logger';
import httpContext = require('express-http-context');
import * as api from '@opentelemetry/api';
import { fetchCurrentSpanCorrelationId } from '../utils';

class MetaDataProvider implements ILoggingMetaDataProvider {
    userId(): string {
        return httpContext.get('userId') || 's2s';
    }

    appName(): string {
        return process.env.NODE_ENV + '-' + process.env.APP_NAME;
    }

    correlationId(): string {
        return fetchCurrentSpanCorrelationId();
    }

    appPackageId(): string {
        return '';
    }

    clientDeviceId(): string {
        return httpContext.get('x-device-id');
    }

    clientPlatform(): string {
        return httpContext.get('x-platform-id'); // Andoid/iOS/Web
    }

    clientSessionId(): string {
        return httpContext.get('x-session-id');
    }

    clientVersion(): string {
        return httpContext.get('x-app-version') || '';
    }

    loggerName(): string {
        return '';
    }

    spanId(): string {
        let current_span = api.trace.getSpan(api.context.active());
        let span_id = current_span?.spanContext().spanId || '';
        return span_id;
    }

    thread(): string {
        return '';
    }

    traceId(): string {
        let current_span = api.trace.getSpan(api.context.active());
        let trace_id = current_span?.spanContext().traceId || '';
        return trace_id;
    }
}

export const logger = new LoggerService(
    process.env.NODE_ENV == 'prod' ? LogLevel.INFO : LogLevel.DEBUG,
    new MetaDataProvider(),
);
