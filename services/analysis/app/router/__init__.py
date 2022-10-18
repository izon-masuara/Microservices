from controller import InputData, Analysis
baseUrl = "/api/v1/analysis"

def router(api):
    api.add_resource(InputData,baseUrl+"/")
    api.add_resource(Analysis,baseUrl+"/<int:id>")