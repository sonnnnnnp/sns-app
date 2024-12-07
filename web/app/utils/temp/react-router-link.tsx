import React from "react";
import { Link, type LinkProps } from "react-router";

interface ReactRouterLinkProps extends Omit<LinkProps, "to"> {
  href: string;
}

const ReactRouterLink = React.forwardRef<
  HTMLAnchorElement,
  ReactRouterLinkProps
>(({ href, ...props }, ref) => {
  // biome-ignore lint/performance/noDelete: <explanation>
  // biome-ignore lint/complexity/useLiteralKeys: <explanation>
  delete props["onClick"];

  return <Link to={href} ref={ref} {...props} />;
});

ReactRouterLink.displayName = "ReactRouterLink";

export default ReactRouterLink;
