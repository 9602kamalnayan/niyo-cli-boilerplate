export enum HTTP_HEADERS_KEY {
    CORRELATION_ID = 'x-correlation-id',
    AUTHORIZATION = 'authorization',
    REQUEST_ID = 'x-request-id',
    IS_TESTING = 'is-testing',
    DEVICE_ID = 'x-device-id',
    APP_VERSION = 'x-app-version',
    PLATFORM = 'x-platform',
}

export enum HTTP_CONTEXT_KEY {
    CORRELATION_ID = 'correlationId',
    APP_USER_ID = 'appUserId',
    USER_DETAILS = 'userDetails',
    USER_ID = 'userId',
    AUTH_TOKEN = 'authToken',
    IS_TESTING = 'isTesting',
    TOKEN_PAYLOAD = 'tokenPayload',
    REQUEST_ID = 'requestId',
    WORKBENCH_PERMISSIONS = 'workbenchPermissions',
}
