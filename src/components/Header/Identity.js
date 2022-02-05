import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faGithub,
  faLinkedin,
  faInstagram,
} from "@fortawesome/free-brands-svg-icons";
import ContactInfo from "./ContactInfo";

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
      <div className="h-full bg-gray-200 overflow-hidden flex flex-row justify-center">
        <div className="flex flex-col justify-center items-center h-full max-w-3xl w-5/6 md:w-11/12">
          <div className="mt-4 mb-auto w-full flex flex-row justify-evenly">
            {/* <h5>Home</h5>
            <h5>About Me</h5>
            <h5>Projects</h5>
            <h5>Blog</h5> */}
          </div>
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
            <span className="font-black"> give my best</span>, improve myself to
            be a better person, and complete my works to a high standard. I am
            flexible, open to new ideas, and constantly looking at how to solve
            current problems with effective solutions by thinking outside of the
            box. I also have leadership experience by leading my campus&apos;
            Google Developer Student Clubs &#40;
            <span className="text-googred font-bold">G</span>
            <span className="text-googblue font-bold">D</span>
            <span className="text-googgreen font-bold">S</span>
            <span className="text-googyellow font-bold">C</span>&#41; and I want
            to make sure that everyone is okay on my team.
          </p>
          {/* <p className="mt-2 w-full text-justify">
            I have freelance experience in Tutoring, Wordpress, and ReactJS with
            Redux.
          </p> */}
          <h3 className="w-full font-bold text-lg mt-4 md:mt-8">Skills:</h3>
          <div className="mt-auto w-full">
            <ContactInfo />
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
