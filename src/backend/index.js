const express = require("express");
const Joi = require("joi");
const fs = require("fs");
const superagent = require("superagent");
const cors = require("cors");
const app = express();
const port = process.env.PORT || 5999;
const gcSecretKey = process.env.GCAPTCHA_SECRET;

const contactFile = "./mycontact.json";

if (!gcSecretKey) {
  console.error("No Google Captcha Secret Key Specified");
  process.exit(1);
}
if (!fs.existsSync(contactFile)) {
  console.error("No Contact file exists");
  process.exit(1);
}

const captchaSchema = Joi.object({
  response: Joi.string().required(),
});

// JSON Decoder
app.use(express.json());

app.use(cors({}));

app.post("/getcontact", (req, res) => {
  const { value, error } = captchaSchema.validate(req.body);
  if (error) return res.sendStatus(400);

  const { response } = value;

  superagent
    .post(`https://www.google.com/recaptcha/api/siteverify`)
    .set({
      Accept: "application/json",
    })
    .query({
      secret: gcSecretKey,
      response,
    })
    .end((err, result) => {
      if (err) return res.sendStatus(400);

      const { success, action, score } = result.body;
      if (!success || action !== "getcontact" || score > 50)
        return res.sendStatus(401);

      res.send(JSON.parse(fs.readFileSync(contactFile)));
    });
});

app.listen(port, () => {
  console.log(`App listening on port ${port}`);
});
