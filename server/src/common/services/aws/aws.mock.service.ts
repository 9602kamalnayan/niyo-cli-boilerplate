import { Injectable } from '@nestjs/common';
import { logger } from 'src/common/helper/logger';

export interface ASMClientConfig {
  region?: string;
  credentials?: any;
  asmRestOptions?: any;
}

export interface AWSCredentialPair {
  accessKeyId: string;
  secretAccessKey: string;
  sessionToken?: string;
}

@Injectable()
export class AwsMockService {
  private secretManger: any;
  private readonly sqs: any;

  constructor() {
    this.sqs = {
      region: 'ap-south-1',
    };
  }

  public createSecretManager(secretManagerConfig: ASMClientConfig): void {
    logger.debug('[Creating secret manager]');
    this.secretManger = {
      username: 'hellops',
      password: 'password',
      SecretString: { clsa_key: 'hellp', clsa_secret: 'paswword' },
    };
    logger.debug('[Secret manager created]');
  }
  public async getSecret(params: any): Promise<any> {
    const data = await this.secretManger;
    if (data?.SecretString) {
      return data.SecretString;
    }
    logger.error('AWS secret-manager secret fetch fail');
    throw new Error('AWS Secrets Manager Empty Data!!!');
  }

  getSqsClient(): any {
    return this.sqs;
  }

  async getKafkaConnectionConfigs(connectionConfig) {
    this.createSecretManager(connectionConfig.secretManager);
    const secrets = await this.getSecret(connectionConfig.SecretId);
    logger.debug('AWS secret-manager configs', secrets);
    if (!secrets.clsa_key || !secrets.clsa_secret) {
      throw new Error('Missing Kafka Connection secret');
    }
    return secrets;
  }

  async getFirebaseConnectionConfigs(connectionConfig) {
    return connectionConfig;
  }
}
