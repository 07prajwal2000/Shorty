import React from "react";
import Navbar from "../components/Navbar";

interface LayoutProps {
	children: React.ReactNode | React.ReactNode[];
}

const DefaultLayout: React.FC<LayoutProps> = ({ children }) => {
	return (
		<div className="container-fluid d-flex flex-column">
			<div className="sticky-top">
				<Navbar />
			</div>
			<main>{children}</main>
		</div>
	);
};

export default DefaultLayout;
