import Navbar from "@/components/navbar";
import UserForm from "@/components/userform/userform";



export default function ProfilePage() {

    return (
        <>
            <Navbar lang="en" />
            <main className="max-w-3xl mx-auto p-2">

                <h1 className="text-xl my-4">Profile Info</h1>
            
                <UserForm />
            </main>
        </>
    )
}