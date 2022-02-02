import * as React from "react";
import Header from "../components/Header/Header";
import { Helmet } from "react-helmet";

const IndexPage = () => {
  return (
    <div>
      <Helmet titleTemplate="Jasoncoding - %s" defaultTitle="Jasoncoding">
        <html lang="en" amp />
        <meta charSet="utf-8" />
        <link rel="canonical" href="https://jasoncoding.com" />
        <meta
          name="description"
          content="Jasoncoding website. My portfolio site"
        />
      </Helmet>
      <Header />
    </div>
  );
};

export default IndexPage;
