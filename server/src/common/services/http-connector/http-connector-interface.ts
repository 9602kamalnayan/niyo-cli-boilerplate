import { OutgoingHttpHeaders } from 'http';
import { AllowedHttpMethod } from '.dto/allowed-http-methods';

export interface IRequestConfig {
  method: AllowedHttpMethod;
  url: string;
  data?: any;
  headers?: OutgoingHttpHeaders;
  additionalData?: any;
  fullResponse?: boolean;
  params?: any;
}
