import { BrowserRouter, Route, Routes } from "react-router-dom";
import IndexPage from "../pages";
import DefaultLayout from "./DefaultLayout";
import Notfound from "../pages/Notfound";

const PageRouter = () => {
	return (
		<BrowserRouter>
			<DefaultLayout>
				<Routes>
					<Route element={<IndexPage />} path="/" />
					<Route element={<Notfound />} path="*" />
				</Routes>
			</DefaultLayout>
		</BrowserRouter>
	);
};

export default PageRouter;
