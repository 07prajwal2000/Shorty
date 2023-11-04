type GlobalStoreType = {
  authSettings: AuthSettings;
};

export type AuthSettings = {
  LoggedIn: boolean;
  SetLoggedIn: (value: boolean) => void;
  Tokens: {Auth: string, Refresh: string};
  SetTokens: (auth: string, refresh: string) => void;
};

export default GlobalStoreType;