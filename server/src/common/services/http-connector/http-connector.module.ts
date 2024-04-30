import { Module, Global } from '@nestjs/common';
import { HttpModule } from '@nestjs/axios';
import { ConnectorService } from './http-connector.service';

@Global()
@Module({
  imports: [
    HttpModule.register({
      timeout: 60000,
      maxRedirects: 5,
    }),
  ],
  providers: [ConnectorService],
  exports: [HttpModule, ConnectorService],
})
export class ConnectorModule {}
