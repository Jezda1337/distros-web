import Link from "next/link";
import { useRouter } from "next/router";
import { Button } from "../ui/Button";

export function AdminHeader() {
  const router = useRouter();

  async function handleLogout() {
    const response = await fetch("/api/v1/auth/logout");
    console.log(await response.json());
    router.push("/");
  }
  return (
    <header className="my-mt-12 my-14 flex items-center justify-between">
      <p className="text-2xl font-bold">Hello Admin</p>

      <div>
        <Link className="hover:underline" href="/">
          Home
        </Link>
        <Button onClick={handleLogout} className="ml-4">
          Logout
        </Button>
      </div>
    </header>
  );
}
