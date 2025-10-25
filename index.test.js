const request = require("supertest");
const app = require("./index");

describe("API Tests", () => {
  describe("GET /api/hello", () => {
    it("正常にレスポンスを返す", async () => {
      const response = await request(app).get("/api/hello");

      expect(response.status).toBe(200);
      expect(response.body).toEqual({ message: "Hello from server!" });
    });

    it("Content-Typeがapplication/jsonである", async () => {
      const response = await request(app).get("/api/hello");

      expect(response.headers["content-type"]).toMatch(/application\/json/);
    });
  });
});
