import React, { useState } from "react";
import Showcase from "./Showcase";

const Header = () => {
  const [skipAnim, setSkipAnim] = useState(false);

  return (
    <div className="w-full h-screen flex flex-col md:flex-row">
      <Showcase
        skipAnim={skipAnim}
        setSkipAnim={() => {
          setSkipAnim(true);
        }}
      />
      <div
        id="idpanel"
        className="bg-gray-100 h-[0%] md:w-[0%] overflow-hidden"
      >
        Coming soon :D
      </div>
    </div>
  );
};

export default Header;
