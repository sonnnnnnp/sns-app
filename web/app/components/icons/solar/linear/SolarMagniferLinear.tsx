import type { SVGProps } from "react";

export function SolarMagniferLinear(props: SVGProps<SVGSVGElement>) {
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
        <circle cx={11.5} cy={11.5} r={9.5}></circle>
        <path strokeLinecap="round" d="M18.5 18.5L22 22"></path>
      </g>
    </svg>
  );
}
