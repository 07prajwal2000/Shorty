import { Github, LogIn } from "lucide-react";
import { Link } from "react-router-dom";

const Navbar = () => {
	return (
		<div className="h-16 grid grid-cols-10 gap-4 bg-slate-900 shadow-blue-700 shadow-sm">
			<a
				title="Shorty - A URL Shortener"
				href="/"
				className="col-span-2 flex flex-row gap-2 items-center px-2"
			>
				<img className="w-12 h-12" src="logo.png" />
				<h2 className="uppercase font-bold text-2xl">Shorty</h2>
			</a>

			<div className="col-span-6 border border-red-400">MAIN ITEMS NAVBAR</div>
			<div className="col-span-2 gap-10 flex justify-center items-center">
				<Link title="Login" to="/login">
					<LogIn className="h-7 w-7" />
				</Link>
				<a title="Star us on Github" href="https://github.com/07prajwal2000/Shorty-The-URL-shortener" target="_blank"><Github className="h-7 w-7" /></a>
			</div>
		</div>
	);
};

export default Navbar;
