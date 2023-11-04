# URLs Controller

##### Base Url: `/api/v1/urls`

- ### Generate URL
  
  - Method: **POST**
  
  - Path: `/generate`
  
  - Description: Generates a new short url from the given original URL.

Request Body Schema:

```typescript
type RequestBody = {
    originalUrl: string;
}
```

Response body Schema:

```typescript
type ResponseBody = {
    original: string; // supplied URL
    modified: string; // any extras added i,e '/'
    shortUrl: string; // Server's domain with hash combined
    hash: string; // generated hash
    statusCode: number; // HTTP Status code
    message: string; // message from the server
};
```

- ### Redirect to Url by ID
  
  - Method: **GET**
  
  - Path: `/:id`. The api accepts a Route parameter. <u>Ex</u>: `/api/v1/urls/s6wzser5`
  
  - Description: Redirects to the URL if the ID is found in Database

Response body Schema:

- If the ID is not found or invalid ID

```typescript
OriginalUrl :string;
StatusCode :number;  
Message :string;
```

- If the ID is found, then it will be automatically redirected.
