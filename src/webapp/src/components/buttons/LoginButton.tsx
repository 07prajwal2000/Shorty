import { useNavigate } from "react-router-dom"
import PageRoutes from "../../Constants/PageRoutes";

const LoginButton = () => {
  const nav = useNavigate();

  function onClick() {
    nav(PageRoutes.Login);
  }
  
  return (
    <button onClick={onClick} className='btn btn-success d-flex justify-content-center align-items-center gap-1' title='Login'>
      <i className="fa-solid fa-right-to-bracket p-0"></i> <span>Login</span>
    </button>
  )
}

export default LoginButton