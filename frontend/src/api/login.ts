export async function loginApi(formdata: { email: string; password: string }) {
  if(formdata.password == 'sujeeta') {
    const user ={
      email: formdata.email,
      token: "FREE TOKEN"
    }
    return user;
  }
  
  throw {
        message: "Failed to Login",
        statusText: "res.statusText",
        status: 400,
      };
}
