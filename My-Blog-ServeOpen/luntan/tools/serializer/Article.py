from rest_framework import serializers

from appone.models.article import Article


class ResourceSerializer(serializers.ModelSerializer):
    class Meta:
        model = Article
        fields = "__all__"

