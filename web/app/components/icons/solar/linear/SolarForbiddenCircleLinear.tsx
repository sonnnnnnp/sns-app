import type { SVGProps } from "react";

export function SolarForbiddenCircleLinear(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={24}
      height={24}
      viewBox="0 0 24 24"
      {...props}
    >
      <g fill="none" stroke="currentColor" strokeWidth={1.5}>
        <path strokeLinecap="round" d="m18.5 5.5l-13 13"></path>
        <circle cx={12} cy={12} r={10}></circle>
      </g>
    </svg>
  );
}
