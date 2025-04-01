import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: "TaskManagementSystem" },
    { name: "Task Management System", content: "Welcome to Task Management System" },
  ];
};

export default function Index() {
  return (
    <div>
        <h1>Hello World</h1>
    </div>
  );
}
