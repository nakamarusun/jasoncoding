import React from "react";
import logo from "../images/logo.svg";

const Header = () => {
  return (
    <div className="w-full h-screen flex flex-row">
      <div className="w-full flex flex-row justify-center bg-black">
        <div className="max-w-4xl w-11/12 h-full flex flex-col justify-center items-center">
          <img
            // // className="max-h-[50%] h-full" // Forgive me, for i have to commit some evilness.
            className="max-h-full filter invert"
            src={logo}
          />
          <h2 className="font-serif text-white text-right w-full text-3xl origin-right scale-x-[70%]">
            COMPUTER SCIENCE STUDENT, FULL-STACK DEVELOPER, DESIGNER
          </h2>
          <div className="material-icons text-white animate-bounce mt-[1rem] mb-[-.25rem]">
            adjust
          </div>
          <button
            style={{
              transition: "all 0.5s cubic-bezier(0.22, 1, 0.36, 1)",
            }}
            className="font-serif border-2 border-gray-50 px-6 hover:px-10 pt-1 pb-2 text-xl font-bold text-white rounded-3xl hover:rounded-lg bg-gradient-to-b from-black to-gray-700"
          >
            PROJECT SHOWCASE
          </button>
        </div>
      </div>
      <div className="bg-gray-100 w-full">Coming soon :D</div>
    </div>
  );
};

export default Header;
