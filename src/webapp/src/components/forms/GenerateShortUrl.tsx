import { useState } from "react";
import UrlServices from "../../api/urlServices";
import { toast } from "react-hot-toast";

const GenerateShortUrl = () => {
	const [originalUrl, setOriginalUrl] = useState("");
	const [loading, setLoading] = useState(false);
	const [shortenUrl, setShortenUrl] = useState("");

	async function onGenerateClicked() {
    if (!UrlServices.ValidateUrl(originalUrl)) {
      toast.error("Enter a valid URL"); 
      return;
    }
    
		setLoading(true);
		setShortenUrl("");
    try {
      const response = await UrlServices.GenerateUrl(originalUrl);
      setShortenUrl(response.shortUrl);
      toast.success("Successfully generated a short url.")
    } catch (error: any) {
      const msg = error.response.data.message || "Server error";
      toast.error(msg); 
    } finally {
      setLoading(false);
    }
	}

	function onCopyClicked() {
		navigator.clipboard.writeText(shortenUrl);
		setLoading(false);
	}

	return (
		<div className="col-12 col-lg-6 border p-3 d-flex flex-column justify-content-center gap-2">
			<h3 className="fw-bold text-primary font-monospace">
				Enter your url here:
			</h3>
			<input
				type="text"
				placeholder="https://domain.com/"
				className="form-control border-3 fw-bold gen-url-input"
				value={originalUrl}
				disabled={loading}
				onChange={(e) => setOriginalUrl(e.target.value)}
			/>
			<button
				disabled={loading}
				onClick={onGenerateClicked}
				className="btn btn-primary w-25 mx-auto mt-2 text-uppercase fw-bold"
			>
				{loading ? (
					<div
						className="spinner-border text-light spinner-border-sm"
						role="status"
					>
					</div>
				) : (
					"Shorten"
				)}
			</button>

			<div className="position-relative">
				{shortenUrl == "" && (
					<div className="position-absolute start-0 end-0 top-0 bottom-0 bg-light z-3" />
				)}
				<p className="text-dark">Here is your shortened url.</p>
				<div className="input-group mb-3">
					<input
						type="text"
						className="form-control fw-bold bg-dark-subtle"
						value={shortenUrl}
						disabled
						readOnly
					/>
					<a
						title="Open Shorten URL in new tab"
            href={shortenUrl}
            target="_blank"
						className="btn btn-success"
						type="button"
					>
						<i className="fa-solid fa-up-right-from-square"></i>
					</a>
					<button
						title="Copy Shorten URL"
						onClick={onCopyClicked}
						className="btn btn-primary"
						type="button"
					>
						<i className="fa-solid fa-clipboard"></i>
					</button>
				</div>
        <div className="border border-2 border-warning bg-warning-subtle rounded-2 d-flex p-2">
        <i className="fa-solid fa-circle-info text-warning fs-4"></i>
        <p className="px-3 my-0 text-warning">Un-Authenticated short urls will last for only 7 days.</p>
        </div>
			</div>
		</div>
	);
};

export default GenerateShortUrl;
