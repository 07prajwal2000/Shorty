import { create } from "zustand";
import GlobalStoreType from "./globalStoreTypes";

const useGlobalStore = create<GlobalStoreType>((set) => ({
	authSettings: {
		LoggedIn: false,
		SetLoggedIn(value) {
			set((p) => ({
				...p,
				authSettings: { ...p.authSettings, LoggedIn: value },
			}));
		},
    Tokens: {Auth: '', Refresh: ''},
    SetTokens(auth, refresh) {
      const token = {
        Auth: auth,
        Refresh: refresh
      };
      set((p) => ({
				...p,
				authSettings: { ...p.authSettings, Tokens: token },
			}));
    },
	},
}));

export default useGlobalStore;