import React, { useState } from "react";
import Identity from "./Identity";
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
      <Identity />
    </div>
  );
};

export default Header;
