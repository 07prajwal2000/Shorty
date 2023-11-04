import { useNavigate } from "react-router-dom";
import PageRoutes from "../../Constants/PageRoutes";

const SignupButton = () => {
  const nav = useNavigate();

  function onClick() {
    nav(PageRoutes.Signup);
  }
  
  return (
    <button onClick={onClick} className='btn btn-info d-flex justify-content-center align-items-center gap-1' title='Login'>
      <i className="fa-solid fa-user-plus"></i> <span>Signup</span>
    </button>
  )
}
export default SignupButton;