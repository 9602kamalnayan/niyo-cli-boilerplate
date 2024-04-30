import { Request, Response, NextFunction } from 'express';
import { logger } from '../helper/logger';
export async function logAndFormatRequest(
  req: Request,
  res: Response,
  next: NextFunction,
) {
  if (req.path != '<RouteBasePrefix>/healthz') {
    logger.info('[REQUEST]', {
      path: req.path,
      method: req.method,
      ip: req.ip,
      protocol: req.protocol,
      hostname: req.hostname,
      headers: req.headers,
      params: req.params,
    });
  }

  //format-request
  next();
}
