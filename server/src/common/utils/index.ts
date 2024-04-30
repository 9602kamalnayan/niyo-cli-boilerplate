import httpContext = require('express-http-context');
import { v4 as uuidv4 } from 'uuid';
import * as api from '@opentelemetry/api';

export function trimStrings(data: any) {
    if (!data) return;

    if (typeof data === 'object') {
        Object.keys(data).forEach((key) => {
            if (typeof data[key] === 'string') {
                data[key] = data[key].trim();
            } else if (Array?.isArray(data[key])) {
                data[key].forEach((element: any) => {
                    trimStrings(element);
                });
            } else if (typeof data[key] === 'object') {
                trimStrings(data[key]);
            }
        });
    }
}

export function isStringifiedJSON(input): boolean {
    try {
        const parsedJSON = JSON.parse(input);
        return typeof parsedJSON === 'object';
    } catch (error) {
        return false;
    }
}

export function createRedisKeyName(key: string) {
    return process.env.REDIS_KEY_PREFIX + key;
}

export async function asyncForEach(array, callback) {
    for (let index = 0; index < array.length; index++) {
        await callback(array[index], index, array);
    }
}

export function isPod(): boolean {
    return process.env.NODE_ENV != 'localdev';
}

export function fetchCurrentSpanCorrelationId() {
    const current_span = api.trace.getSpan(api.context.active());
    const correlationIdKey = api.createContextKey('correlationId');
    const correlationId =
        api?.context?.active()?.getValue(correlationIdKey) ||
        current_span?.['attributes']['correlationId'];
    return correlationId || httpContext.get('x-correlation-id') || uuidv4();
}
