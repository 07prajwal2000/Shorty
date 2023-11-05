import { useEffect } from "react";
import { useParams } from "react-router-dom";
import UrlService from "../api/urlServices";
import useSWR from "swr";
import Notfound from "../components/sections/Notfound";
import DefaultLayout from "../defaults/DefaultLayout";

const Redirect = () => {
	const { id } = useParams();
	const { data, error, isLoading } = useSWR(
		id,
		(id) => UrlService.GetShortUrl(id),
		{ shouldRetryOnError: false }
	);

	useEffect(() => {
		if (!isLoading && !error) {
			window.location.href = data!.originalUrl;
		}
	}, [isLoading]);

	return (
		<div>
			{isLoading && <LoadingSection />}
			{error && !isLoading && <ErrorSection />}
		</div>
	);
};

function LoadingSection() {
	return <div className="container d-flex align-items-center justify-content-center p-5">
    <div className="spinner-grow text-info" role="status">
  <span className="sr-only">Loading...</span>
</div>
  </div>;
}

function ErrorSection() {
	return (
		<DefaultLayout>
			<Notfound />
		</DefaultLayout>
	);
}

export default Redirect;
