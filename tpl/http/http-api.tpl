{
  "openapi": "3.0.1",
  "info": {
    "title": "",
    "description": "",
    "version": "1.0.0"
  },
  "tags": [
    {
      "name": "{{.xfilename}}"
    }
  ],
  "paths": {
    "/{{.modelname}}/{{.xfilename}}/{id}": {
      "delete": {
        "summary": "删除{{.xfilename}}",
        "x-apifox-folder": "{{.xfilename}}",
        "x-apifox-status": "developing",
        "deprecated": false,
        "description": "",
        "tags": [
          "{{.xfilename}}"
        ],
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "{{.idCommand}}",
            "required": true,
            "example": "",
            "schema": {
              "type": "{{.idType}}"
            }
          }
        ],
        "responses": {

        },
        "x-run-in-apifox": "",
        "security": [
          {
            "bearer": []
          }
        ]
      }
    },
    "/{{.modelname}}/{{.xfilename}}Info": {
      "get": {
        "summary": "根据ID查询{{.xfilename}}",
        "x-apifox-folder": "{{.xfilename}}",
        "x-apifox-status": "developing",
        "deprecated": false,
        "description": "",
        "tags": [
          "{{.xfilename}}"
        ],
        "parameters": [
          {
            "name": "id",
            "in": "query",
            "description": "{{.idCommand}}",
            "required": false,
            "example": "",
            "schema": {
              "type": "{{.idType}}"
            }
          }
        ],
        "responses": {

        },
        "x-run-in-apifox": "",
        "security": [
          {
            "bearer": []
          }
        ]
      }
    },
    "/{{.modelname}}/{{.xfilename}}": {
      "post": {
        "summary": "添加{{.xfilename}}",
        "x-apifox-folder": "{{.xfilename}}",
        "x-apifox-status": "developing",
        "deprecated": false,
        "description": "",
        "tags": [
          "{{.xfilename}}"
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                {{.Add}}
                }
              }
            }
          }
        },
         "responses": {

         },
        "x-run-in-apifox": "",
        "security": [
          {
            "bearer": []
          }
        ]
      },
      "get": {
        "summary": "分页查询{{.xfilename}}",
        "x-apifox-folder": "{{.xfilename}}",
        "x-apifox-status": "developing",
        "deprecated": false,
        "description": "",
        "tags": [
          "{{.xfilename}}"
        ],
        "parameters": [
          {{.Query}}
        ],
         "responses": {

         },
        "x-run-in-apifox": "",
        "security": [
          {
            "bearer": []
          }
        ]
      },
      "put": {
        "summary": "修改{{.xfilename}}",
        "x-apifox-folder": "{{.xfilename}}",
        "x-apifox-status": "developing",
        "deprecated": false,
        "description": "",
        "tags": [
          "{{.xfilename}}"
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                 {{.Update}}
                },
                "required": [
                  "id"
                ]
              }
            }
          }
        },
        "responses": {

        },
        "x-run-in-apifox": "",
        "security": [
          {
            "bearer": []
          }
        ]
      }
    }
  },
  "components": {
    "schemas": {}
  }
}