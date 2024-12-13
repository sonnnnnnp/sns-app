import type { SVGProps } from "react";

export function SolarUserLinear(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      aria-hidden="true"
      xmlns="http://www.w3.org/2000/svg"
      width={24}
      height={24}
      viewBox="0 0 24 24"
      {...props}
    >
      <g fill="none" stroke="currentColor" strokeWidth={2}>
        <circle cx={12} cy={6} r={4}></circle>
        <path d="M20 17.5c0 2.485 0 4.5-8 4.5s-8-2.015-8-4.5S7.582 13 12 13s8 2.015 8 4.5Z"></path>
      </g>
    </svg>
  );
}
