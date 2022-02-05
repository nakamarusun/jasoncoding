import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faGithub,
  faLinkedin,
  faInstagram,
} from "@fortawesome/free-brands-svg-icons";
import ContactInfo from "./ContactInfo";
import "./Identity.css";
import { menus } from "../../menu";

const socials = [
  {
    icon: faGithub,
    link: "https://github.com/nakamarusun",
  },
  {
    icon: faLinkedin,
    link: "https://linkedin.com/in/jasoncoding/",
  },
  {
    icon: faInstagram,
    link: "https://instagram.com/jason.christianh",
  },
];

const Identity = () => {
  return (
    <div id="idpanel" className="h-[0%] md:w-[0%]">
      <div className="h-full overflow-hidden flex flex-row justify-center">
        <div className="flex flex-col justify-center items-center h-full max-w-3xl w-5/6 md:w-11/12">
          <div className="flex-1 flex flex-col justify-center md:mt-16">
            <h2 className="text-gray-600 font-medium text-xl w-full leading-none ml-8 md:ml-0 mt-8 md:mt-0">
              Hello,
            </h2>
            <h1 className="text-4xl font-black w-full ml-8 md:ml-0">
              I&apos;m&nbsp;
              <span className="group relative">
                <span className="absolute text-sm text-gray-800 w-full text-center ease-out duration-200 opacity-0 group-hover:opacity-100">
                  Call me
                </span>
                <span className="text-purple-700 ease-out duration-200 group-hover:translate-y-2 inline-block">
                  Jason&nbsp;
                </span>
              </span>
              Christian.
            </h1>
            <p className="mt-4 md:mt-8 w-full">
              Thank you for stopping by my site!
            </p>
            <p className="mt-2 w-full text-justify">
              I am an 🇮🇩 developer highly motivated to
              <span className="font-black"> give my best</span>, improve myself
              to be a better person, and complete my works to a high standard. I
              am flexible, open to new ideas, and constantly looking at how to
              solve current problems with effective solutions by thinking
              outside of the box. I also have leadership experience by leading
              my campus&apos; Google Developer Student Clubs &#40;
              <span className="text-googred font-bold">G</span>
              <span className="text-googblue font-bold">D</span>
              <span className="text-googgreen font-bold">S</span>
              <span className="text-googyellow font-bold">C</span>&#41; and I
              want to make sure that everyone is okay on my team.
            </p>
            {/* <p className="mt-2 w-full text-justify">
              I have freelance experience in Tutoring, Wordpress, and ReactJS with
              Redux.
            </p> */}
            <h3 className="w-full font-bold text-lg mt-4 md:mt-8">Skills:</h3>
            <ContactInfo className="mt-4" />
          </div>
          <div className="w-full">
            <div className="hidden mb-5 w-full md:flex flex-row justify-evenly">
              {menus.map(({ name, url }, i) => (
                <a
                  className="bg-white rounded-md px-3 py-0.5 shadow-md text-sm font-medium transition-colors duration-25 underline decoration-transparent hover:bg-gray-100 hover:decoration-gray-500"
                  key={i}
                  href={url}
                >
                  {name}
                </a>
              ))}
            </div>
            <div className="w-full flex flex-row justify-evenly mb-6">
              {socials.map(({ icon, link }, i) => (
                <a target="_blank" href={link} key={i} rel="noreferrer">
                  <FontAwesomeIcon
                    className="text-purple-900 hover:text-fuchsia-700"
                    icon={icon}
                    size="2x"
                  />
                </a>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Identity;
// TODO: Tech stack animation flying everywhere
