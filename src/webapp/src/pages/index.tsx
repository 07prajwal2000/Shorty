import GenerateShortUrl from "../components/forms/GenerateShortUrl";
import Hero1 from "../components/sections/Hero1";

const IndexPage = () => {
	return (
		<div className="container">
			{/* Hero 1 section */}
			<div className="row my-4">
				<GenerateShortUrl />
				<Hero1 />
			</div>
		</div>
	);
};

export default IndexPage;
