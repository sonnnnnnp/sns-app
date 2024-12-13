import type { SVGProps } from "react";

export function SolarMicrophone2Linear(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={24}
      height={24}
      viewBox="0 0 24 24"
      {...props}
    >
      <g fill="none" stroke="currentColor" strokeWidth={2}>
        <path d="M7 8a5 5 0 0 1 10 0v3a5 5 0 0 1-10 0z"></path>
        <path
          strokeLinecap="round"
          d="M13.5 8H17m-3.5 3H17M7 8h2m-2 3h2m11-1v1a8 8 0 1 1-16 0v-1m8 9v3"
        ></path>
      </g>
    </svg>
  );
}
