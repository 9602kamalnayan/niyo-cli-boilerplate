import { Module } from '@nestjs/common';
import { AwsService } from './aws.service';
import { AwsMockService } from './aws.mock.service';

@Module({
  providers: [
    {
      provide: AwsService,
      useClass:
        process.env.NODE_ENV !== 'localdev' ? AwsService : AwsMockService,
    },
  ],
  exports: [AwsService],
})
export class AwsModule {}
