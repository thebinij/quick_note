import {
  Form,
  redirect,
  useActionData,
  useLoaderData,
  useNavigation,
} from "react-router-dom";
import { loginApi } from "../api/login";

export function loader({ request }: any) {
  return new URL(request.url).searchParams.get("message");
}

export async function action({ request }: any) {
  const formData = await request.formData();
  const email = formData.get("email");
  const password = formData.get("password");
  try {
    const data = await loginApi({ email, password });
    if (data?.token) {
      localStorage.setItem("USER_CRED", data.token);
      return redirect("/profile");
    }
  } catch (err: any) {
    return err?.message;
  }
}

function Login() {
  const navigation = useNavigation();
  const message = useLoaderData() as string;
  const errorMessage = useActionData() as string;
  const inputStyle = {
    width: "250px",
    height: "40px",
    margin: "10px",
    padding: "5px",
  };

  return (
    <div>
      {message && <h3>{message}</h3>}
      {errorMessage && <h3>{errorMessage}</h3>}

      <Form
        method="post"
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
        }}
        replace
      >
        <input
          name="email"
          type="text"
          placeholder="Email"
          style={inputStyle}
        />
        <input
          name="password"
          type="password"
          placeholder="Password"
          style={inputStyle}
        />

        <button
          style={{
            width: "250px",
            height: "40px",
            margin: "10px",
            backgroundColor:
              navigation.state == "submitting" ? "grey" : "#4CAF50",
            color: "white",
            border: "none",
            borderRadius: "5px",
            cursor: navigation.state == "submitting" ? "none" : "pointer",
          }}
          disabled={navigation.state == "submitting"}
        >
          {navigation.state == "submitting" ? "Logging in..." : "Login"}
        </button>
      </Form>
    </div>
  );
}

export default Login;
