import { Credentials, SecretsManager as ASM, SQS } from 'aws-sdk';
import { Injectable } from '@nestjs/common';
import { logger } from 'src/common/helper/logger';

export interface ASMClientConfig {
  region?: string;
  credentials?: AWSCredentialPair | Credentials;
  asmRestOptions?: ASM.ClientConfiguration;
}

export interface AWSCredentialPair {
  accessKeyId: string;
  secretAccessKey: string;
  sessionToken?: string;
}

@Injectable()
export class AwsService {
  private secretManger: ASM;
  private readonly sqs: SQS;

  constructor() {
    this.sqs = new SQS({
      region: 'ap-south-1',
    });
  }

  private createSecretManager(secretManagerConfig: ASMClientConfig): void {
    try {
      logger.info('[Creating secret manager]');
      const clientConfig: ASM.ClientConfiguration = {
        ...secretManagerConfig?.asmRestOptions,
        ...{
          region: secretManagerConfig.region,
          credentials: null,
        },
      };
      this.secretManger = new ASM(clientConfig);
      logger.info('[Secret manager created]');
    } catch (err) {
      logger.error('ERROR IN createSecretManager', err);
    }
  }
  private async getSecret(
    params: ASM.Types.GetSecretValueRequest,
  ): Promise<any> {
    try {
      const data = await this.secretManger.getSecretValue(params).promise();
      if (data?.SecretString) {
        return JSON.parse(data.SecretString);
      }
      logger.info('AWS secret-manager secret fetch fail');
      return {};
    } catch (err) {
      logger.error('ERROR IN getSecret', err);
    }
  }

  getSqsClient(): SQS {
    return this.sqs;
  }

  async getKafkaConnectionConfigs(connectionConfig) {
    try {
      this.createSecretManager(connectionConfig.secretManagerConfig);
      logger.info('Fetching Kafka Secrets');
      const secrets = await this.getSecret({
        SecretId: connectionConfig.SecretId,
      });
      if (!secrets.clsa_key || !secrets.clsa_secret) {
        logger.error('Missing secrets for Kafka Connection', secrets);
        throw new Error('Missing Kafka Connection secret');
      }
      logger.info('Fetched Kafka Secrets');
      logger.debug(secrets);
      return secrets;
    } catch (err) {
      logger.error('ERROR IN getKafkaConnectionConfigs', err);
    }
  }
}
