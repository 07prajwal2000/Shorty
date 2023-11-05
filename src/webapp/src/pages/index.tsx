import GenerateShortUrl from "../components/forms/GenerateShortUrl";
import Hero1 from "../components/sections/Hero1";
import DefaultLayout from "../defaults/DefaultLayout";

const IndexPage = () => {
	return (
		<DefaultLayout>

		<div className="container">
			{/* Hero 1 section */}
			<div className="row my-4">
				<GenerateShortUrl />
				<Hero1 />
			</div>
		</div>
		</DefaultLayout>
	);
};

export default IndexPage;
