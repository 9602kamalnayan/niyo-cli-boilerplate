import { Injectable, NestMiddleware } from '@nestjs/common';
import httpContext = require('express-http-context');
import { HTTP_CONTEXT_KEY, HTTP_HEADERS_KEY } from '../constants';
import api from '@opentelemetry/api';

@Injectable()
export class setHttpContextMiddleware implements NestMiddleware {
  use(req: any, res: any, next: () => void) {
    let current_span = api.trace.getSpan(api.context.active());
    let traceId = current_span?.spanContext()?.traceId || '';

    let correlationId =
      req.headers[HTTP_HEADERS_KEY.CORRELATION_ID] || `x2-core-arab-${traceId}`;
    res.setHeader(HTTP_HEADERS_KEY.CORRELATION_ID, correlationId);

    httpContext.ns.bindEmitter(req);
    httpContext.ns.bindEmitter(res);
    httpContext.set(HTTP_CONTEXT_KEY.CORRELATION_ID, correlationId);
    httpContext.set(HTTP_CONTEXT_KEY.REQUEST_ID, `x2-core-arab-${traceId}`);
    next();
  }
}
