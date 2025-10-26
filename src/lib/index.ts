// place files you want to import through the `$lib` alias in this folder.

type Link = {
  url: string;
  name: string;
}

export const SITE_NAME = "MedWorld";
export const SITE_DESCRIPTION =
  "MedWorld - Your one-stop solution for all medical needs.";

export const links: Link[] = [
  {
        name: "Features",
        url: "/features"
  },
  {
        name: "About",
        url: "/about"
  },
  {
        name: "Pricing",
        url: "/pricing"
  },
  {
        name: "Contact",
        url: "/contact"
  },
]
