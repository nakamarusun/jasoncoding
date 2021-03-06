import * as React from "react";
import Header from "../components/Header/Header";
import { Helmet } from "react-helmet";
import { GoogleReCaptchaProvider } from "react-google-recaptcha-v3";
import Sidebar from "../components/Sidebar";
import splash from "../images/splash.jpg";

const IndexPage = () => {
  return (
    <GoogleReCaptchaProvider
      reCaptchaKey="6LdGuFkeAAAAAFJRKFLPVB5Cd51jY0R1GKpsCZnL"
      scriptProps={{ defer: true, async: true }}
    >
      <Helmet
        titleTemplate="Jasoncoding - %s"
        defaultTitle="Jasoncoding - Full Stack Engineer + Designer"
      >
        <html lang="en" amp />
        <meta charSet="utf-8" />
        <link rel="canonical" href="https://jasoncoding.com" />
        <meta
          name="description"
          content="Jason Christian H. - Developer, Designer, Computer Engineer. Welcome to my blog and portfolio."
        />
        {/* Og graphs */}
        <meta property="og:url" content="https://jasoncoding.com/" />
        <meta property="og:type" content="website" />
        <meta
          property="og:title"
          content="Jasoncoding - Full Stack Engineer + Designer"
        />
        <meta
          property="og:description"
          content="Hello! My name is Jason Christian Hailianto. This is my portfolio and blogging website about my projects and everything else."
        />
        <meta name="image" property="og:image" content={splash} />
        <meta name="author" content="Jason Christian" />
        {/* Twitter graphs */}
        <meta name="twitter:card" content="summary_large_image" />
        <meta property="twitter:domain" content="jasoncoding.com" />
        <meta property="twitter:url" content="https://jasoncoding.com/" />
        <meta
          name="twitter:title"
          content="Jasoncoding - Full Stack Engineer + Designer"
        />
        <meta
          name="twitter:description"
          content="Hello! My name is Jason Christian Hailianto. This is my portfolio and blogging website about my projects and everything else."
        />
        <meta name="twitter:image" content={splash} />
      </Helmet>
      <Sidebar />
      <Header />
    </GoogleReCaptchaProvider>
  );
};

export default IndexPage;
