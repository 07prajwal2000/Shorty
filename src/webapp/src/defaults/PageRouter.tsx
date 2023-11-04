import { BrowserRouter, Route, Routes } from "react-router-dom";
import IndexPage from "../pages";
import DefaultLayout from "./DefaultLayout";
import Notfound from "../pages/Notfound";
import PageRoutes from "../Constants/PageRoutes";
import Auth from "../pages/Auth";

const PageRouter = () => {
	return (
		<BrowserRouter>
			<DefaultLayout>
				<Routes>
					<Route element={<IndexPage />} path={PageRoutes.Home} />
					<Route element={<Auth />} path={PageRoutes.Login} />
					<Route element={<Notfound />} path="*" />
				</Routes>
			</DefaultLayout>
		</BrowserRouter>
	);
};

export default PageRouter;
