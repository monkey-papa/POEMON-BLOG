# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.label import Label


class UpdateLabelView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            id = data['id']
            sortId = data['sortId']
            labelDescription = data['labelDescription']
            labelName = data['labelName']
            if id and sortId and labelDescription and labelName:
                Label.objects.filter(id=id).update(sort_id=sortId, label_description=labelDescription,
                                                   label_name=labelName)
                return Response({
                    'result': "success",
                })
            else:
                return Response({
                    'failure': "exists null",
                })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })
