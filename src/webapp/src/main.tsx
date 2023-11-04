import ReactDOM from "react-dom/client";
import "./index.css";
import PageRouter from "./defaults/PageRouter";
import { Toaster } from "react-hot-toast";

ReactDOM.createRoot(document.getElementById("root")!).render(
	<>
		<PageRouter />
		<Toaster
			toastOptions={{
				duration: 3750,
				position: "top-right",
			}}
		/>
	</>
);
