import useGlobalStore from "../../store/global";

const Notfound = () => {
	const {
		authSettings: { LoggedIn },
	} = useGlobalStore();

	return (
		<div className="container d-flex flex-column align-items-center justify-content-center">
			<div>
				<h1
					className="font-monospace text-danger"
					style={{ fontSize: "11rem" }}
				>
					404
				</h1>
				<p className="fs-4">OOPS! Nothing was found.</p>
				<p>The page you are looking for is unavailable.</p>
			</div>
			{!LoggedIn ? <BackToHomeBtn /> : <BackToDashboardBtn />}
		</div>
	);
};

const BackToHomeBtn = () => {
	return (
		<button className="btn btn-info">
			<i className="fa-solid fa-house"></i> <span>Back to Home</span>
		</button>
	);
};

const BackToDashboardBtn = () => {
	return (
		<button className="btn btn-success">
			<i className="fa-solid fa-gauge"></i> <span>Back to Dashboard</span>
		</button>
	);
};

export default Notfound;
