import { BrowserRouter, Route, Routes } from "react-router-dom";
import IndexPage from "../pages";
import PageRoutes from "../Constants/PageRoutes";
import Auth from "../pages/Auth";
import Redirect from "../pages/Redirect";

const PageRouter = () => {
	return (
		<BrowserRouter>
				<Routes>
					<Route element={<IndexPage />} path={PageRoutes.Home} />
					<Route element={<Auth />} path={PageRoutes.Login} />
					<Route element={<Redirect />} path="/:id" />
				</Routes>
		</BrowserRouter>
	);
};

export default PageRouter;
