import ReactDOM from "react-dom/client";
import "./index.css";
import PageRouter from "./defaults/PageRouter";
import { ThemeProvider } from "./@/components/ThemeProvider";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <ThemeProvider>
    <PageRouter />
  </ThemeProvider>
);
