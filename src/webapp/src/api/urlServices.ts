import { GenerateUrlResponse, GetShortUrlResponse } from "../models/url";
import { GetFetcher, PostFetcher } from "./fetcher";

export default class UrlService {
  private static BASE_URL = "http://localhost:5556";
  private static URL_REGEX = /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)/;
  
  private static urls = {
    generate: this.BASE_URL + '/api/v1/urls/generate',
    redirectUrl: this.BASE_URL + '/api/v1/urls/redirect/', // :id is required
  };

  public static async GenerateUrl(originalUrl: string) {
    const result = await PostFetcher(this.urls.generate, {
      originalUrl
    });
    return result.data as GenerateUrlResponse;
  }

  public static async GetShortUrl(id: string) {
    const result = await GetFetcher(this.urls.redirectUrl + id);
    return result.data as GetShortUrlResponse;
  }

  public static ValidateUrl(value: string): boolean {
    return this.URL_REGEX.test(value);
  }
}