import React from "react";

const Identity = () => {
  return (
    <div id="idpanel" className="h-[0%] md:w-[0%]">
      <div className="h-full bg-gray-200 overflow-hidden flex flex-row justify-center">
        <div className="flex flex-col justify-center items-center h-full max-w-3xl w-11/12">
          <h2 className="text-gray-600 font-medium text-xl w-full leading-none ml-8 mt-8 md:mt-0">
            Hello,
          </h2>
          <h1 className="text-4xl font-black w-full ml-8">
            I&apos;m <span className="text-purple-700">Jason</span> Christian.
          </h1>
          <p className="mt-4 md:mt-8 w-full">
            Thank you for stopping by into my site!
          </p>
          <p className="mt-2 w-full text-justify">
            I am a developer highly motivated to
            <span className="font-black"> give my best</span>, improve myself to
            be a better person, and complete my works to a high standard. I am
            flexible, open to new ideas, and constantly looking at how to solve
            current problems with effective solutions by thinking outside of the
            box. I also have leadership experience by leading my campus&apos;
            <span className="text-googred font-bold"> Google</span>
            <span className="text-googblue font-bold"> Developer</span>
            <span className="text-googgreen font-bold"> Student</span>
            <span className="text-googyellow font-bold"> Clubs</span> and I want
            to make sure that everyone is fine on my team.
          </p>
          <p className="mt-2 w-full text-justify">
            I have freelance experience in Tutoring, Wordpress, and ReactJS with
            Redux.
          </p>
        </div>
      </div>
    </div>
  );
};

export default Identity;
// TODO: Tech stack animation flying everywhere
