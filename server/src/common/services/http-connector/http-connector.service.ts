import {
  HttpStatus,
  Injectable,
  InternalServerErrorException,
} from '@nestjs/common';
import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';
import { logger } from '../../helper/logger';

const axiosTimeoutSeconds = process.env.AXIOS_TIMEOUT_SECONDS || 50;
@Injectable()
export class ConnectorService {
  public async request(requestObject: AxiosRequestConfig): Promise<{
    responseCode: number;
    responseData: any;
    responseHeaders: any;
  }> {
    let startTime = null;
    let endTime = null;
    try {
      logger.debug('[NETWORK_CALL_REQUEST] ', {
        url: requestObject.url,
        method: requestObject.method,
        data: requestObject.data,
        headers: requestObject.headers,
      });

      const { url, method, data, headers, params } = requestObject;
      // Handle single and multi-valued query parameters
      const queryParams = new URLSearchParams(params).toString();

      const finalUrl = queryParams.length ? `${url}?${queryParams}` : url;
      const finalRequestConfig = {
        url: finalUrl,
        method,
        data,
        headers,
        timeout: Number(axiosTimeoutSeconds) * 1000, //timeout in ms
      };
      startTime = new Date().getTime();
      logger.info('[ATTEMPTING_NETWORK_CALL] ', {
        url: finalRequestConfig.url,
        method: finalRequestConfig.method,
        data: finalRequestConfig.data,
      });
      const response: AxiosResponse = await axios.request(finalRequestConfig);
      endTime = new Date().getTime();
      logger.info(`[TIME_ELAPSED] ${endTime - startTime} ms`);
      let responseToSend = {
        responseCode: response.status,
        responseData: response.data,
        responseHeaders: response.headers || null,
      };
      logger.info(`[RECEIVED_2XX_RESPONSE] `, responseToSend);
      return responseToSend;
    } catch (error) {
      endTime = new Date().getTime();
      logger.info(`[TIME_ELAPSED] ${endTime - startTime} ms`);
      if (error.response) {
        let responseToSend = {
          responseCode: error.response.status,
          responseData: error.response.data,
          responseHeaders: error.response.headers || null,
        };
        logger.info(`[RECEIVED_NON_2XX_RESPONSE] `, responseToSend);
        return responseToSend;
      } else if (error.request) {
        logger.info(`[CANNOT_GET_RESPONSE] AxiosErrorCode -> ${error.code}`);
        logger.error(error.message);
        return {
          responseCode: HttpStatus.GATEWAY_TIMEOUT,
          responseData: null,
          responseHeaders: null,
        };
      } else {
        // Something happened in setting up the request that triggered an Error
        logger.error(error.message);
        throw new InternalServerErrorException('Technical Error. ');
      }
    }
  }
}
