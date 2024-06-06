import tracer from "./common/utils/metrics-monitoring";
import "./common/utils/dd-tracer";
import { NestFactory } from "@nestjs/core";
import {
  SwaggerModule,
  DocumentBuilder,
  SwaggerDocumentOptions,
} from "@nestjs/swagger";
import { AppModule } from "./app.module";
import { ValidationPipe } from "@nestjs/common";
import * as api from "@opentelemetry/api";
import { AsyncHooksContextManager } from "@opentelemetry/context-async-hooks";
import httpContext = require("express-http-context");
import { logAndFormatRequest } from "./common/middleware/incoming-request-logger";

const contextManager = new AsyncHooksContextManager();
contextManager.enable();
api.context.setGlobalContextManager(contextManager);

function setUpSwagger(app) {
  const config = new DocumentBuilder()
    .setTitle("$ServiceName")
    .setDescription("$ServiceName : API Documentation")
    .setVersion("1.0")
    .build();
  const options: SwaggerDocumentOptions = {
    ignoreGlobalPrefix: true,
  };
  const document = SwaggerModule.createDocument(app, config, options);
  SwaggerModule.setup(
    `$ServiceName/${process.env.NODE_ENV || "dev"}/<RoutePath>`,
    app,
    document
  );
}

async function bootstrap() {
  await tracer.start();
  const app = await NestFactory.create(AppModule);
  app.useGlobalPipes(new ValidationPipe());
  app.setGlobalPrefix("$ServiceName");
  app.use(logAndFormatRequest);
  app.use(httpContext.middleware);
  if (["localdev", "uat"].includes(process.env.NODE_ENV)) {
    setUpSwagger(app);
  }
  await app.listen(3000);
}
bootstrap();
