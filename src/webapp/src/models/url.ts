export type GenerateUrlResponse = {
  original: string;
  modified: string;
  shortUrl: string;
  hash: string; 
  statusCode: number;
  message: string;
};

export type GetShortUrlResponse = {
  originalUrl: string;
  statusCode: number;
  message: string;
};