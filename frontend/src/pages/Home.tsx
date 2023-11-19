import { useLoaderData } from "react-router-dom";

export function loader() {
   return "Data from Server is here";
}

function Home() {
  const data = useLoaderData();
  console.log(data);
  return (
    <main>
      <h2>I am at HOME page!!!!</h2>
    </main>
  )
}

export default Home;
