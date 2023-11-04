import { Link } from "react-router-dom"
import LoginButton from "./buttons/LoginButton"
import SignupButton from "./buttons/SignupButton"
import PageRoutes from "../Constants/PageRoutes"

const Navbar = () => {
  return (
    <div className="border border-3 mt-2 rounded-2 z-3 bg-light row align-items-center shadow shadow-lg mb-1 mx-auto w-100 justify-content-center py-2 px-3">
      <div className="col-3 text-center">
        <Link className="text-decoration-none" to={PageRoutes.Home}><h4 className="fw-bold d-inline text-dark my-auto text-lowercase cursor-pointer">Shorty</h4></Link>
      </div>
      <div className="col-6 border py-2"></div>
      <div className="col-2 ms-auto d-flex flex-row gap-2">
        <SignupButton />
        <LoginButton />
      </div>
    </div>
  )
}

export default Navbar