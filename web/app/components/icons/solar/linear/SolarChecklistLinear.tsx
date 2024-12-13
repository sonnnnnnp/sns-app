import type { SVGProps } from "react";

export function SolarChecklistLinear(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={24}
      height={24}
      viewBox="0 0 24 24"
      {...props}
    >
      <g
        fill="none"
        stroke="currentColor"
        strokeLinecap="round"
        strokeWidth={2}
      >
        <path
          strokeLinejoin="round"
          d="M2 5.5L3.214 7L7.5 3M2 12.5L3.214 14L7.5 10M2 19.5L3.214 21L7.5 17"
        ></path>
        <path d="M22 19H12m10-7H12m10-7H12"></path>
      </g>
    </svg>
  );
}
